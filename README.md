# greedyGamePalash
Backend Assignment | GreedyGame

![design](https://i.imgur.com/5eDfsXX.jpg)
Basic Architecture:

Http Module insert/fetch request are pushed to RequestProcessor channel. Insert call is async and returns 200 imideatlly. Fetch call is sync and waits for response from fetchProcessor to return back. 

RequestProcessor fans out the request based on type(insert or fetch) to InsertProcessor or FetchProcessor


KEY POINTS:
1) RequestProcessor also sends all the request of insert type to InsertLogProcessor as well. InsertLogProcessor writes down the insert queries to log, so that they can be used to reconstruct the Data collection state incase of system failure.

2) The updates to the data collection state are written in thread safe manner. All updates to counter values are done atomicaly and rare case of duplication insertion of any new node is prevented by using mutexes.

3) The Insert request is processed asyncrhronously as requested.

4) Maintainability: Adding new levels like "States" and or "Cities" can be easyly handled by adding their level/priority to the LevelOrder enum.

5) Maintainability: Adding new functionality like remove node can also be easily added in thread safe manner by adding new processor (like insertProcessor). 


IMPROVEMENTS POSSIBLE:
1) Request validation
2) Graceful termination by adding quit channel for all processors
