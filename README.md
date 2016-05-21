# GoAgg
Golang service to fetch and store ATOM feeds for FeedGopher.

Keeps track of a user's ATOM URIs, updates feed information on login and periodically while user is logged in, parses ATOM XML and stores as a document

--note: PriorityQueue implemented with go package container/heap to store list of users feeds to update, prioritized on logged-in status.

Database (https://github.com/lib/pq) structure:
<SQL table builder image>
<Need to phyically draw these>

- User
	- credentials
	- list of feed groups?
	- 

- Feed_FeedGroup
 	- feedID
 	- FeedGroupID

TODO How Feeds are structured and how they can be compiled into a time series
TODO Care for duplication of feed items, how long do we hold them? Until all users have 'read' said item or beyond _ days old?

- Feed
	- URI
	- Items <DOCUMENT>
	- 

- FeedGroup

Mostly written as a separate service from FeedGopher to exercise different technologies.

