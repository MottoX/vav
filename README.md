# Viewed-Also-Viewed Recommendation
Viewed-Also-Viewed is frequently applied in recommender systems. To generate viewed-also-viewed data, we need to calculate `degree` 
between different items. Two items are said to have a degree of `K` if they are viewed by `K` users. An interval threshold is needed
to ensure that different views for each user are related to a certain extent. This project uses Hadoop Streaming, Redis and Go
programming language to implement the data generation process as well as a web interface for queries.

## Environment
The following environment is for the implementation in this project.
```
Go 1.12
Hadoop Streaming
Docker
```

## Map Reduce Process
### Round 1
The main goal for round 1 mapper is to filter and convert the data into the format as follows.
```
userId timeStamp itemId
```
During the process, non-view records will be filtered out by checking the operation field.

With Hadoop's internal mechanism, we let the `timeStamp` be the secondary sorting key while `userId` is the partition key.
The benefit of doing so is that a slide window, basically a FIFO queue, can be used to remove all the items that has a timestamp less than the current one by more than a given threshold.

For reducer, the output format is as follows.
```
itemId1 itemId2
```
Each record indicates that there exists one user contributing 1 degree to the two viewed items.
Duplication reduction is applied here. For example, a user clicked item 1 and then clicked item 2, and clicked item 1 once again. In such cases, 
item 1 should increase only 1 degree with item 2 and vice versa.

### Round 2
In the subsequent round 2, we only need a reducer which uses a map for each item to calculate degrees with other items.
For each item, its associated items are sorted in the descending order of degree. The format is as follows.
```
itemId    itemId1    degree1 itemId2    degree2 ...
```
where `itemId1` and `itemId2` represent the indices of the associated items.

### Redis
Each item record is stored in Redis using a list as follows.
```
itemId [itemId1-degree1, itemId2-degree2, ...]
```
where the key is item id.

## Deployment
To run Redis and web service, simply execute the following command.
```
docker-compose up
```
Run `store_result.sh` in `db` folder to store the viewed-also-viewed result into Redis.
Use the following URL to query results.
```
http://localhost:8080/vav?item_id=xx&min_degree=yy
```