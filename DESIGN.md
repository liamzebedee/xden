# Design
The design requirements are intended to guide the development of this project. They are not intended to serve as a reference to the design of XDen, merely a brief overview of what must be accomplished. 

### Data storage structure
* key-value data store
* arbitrary keys 
* mutable values

### Overlay network
#### Structure
* Consists of contacts routing messages to distribute objects. A contact is made up:
 - a node (IPv6 address + port)
 - an ID (of type Key) 
 - an entry contact
* Contacts enter the network via an **entry contact** (bootstrapping)
* In bootstrapping, the contact gathers information about other nodes in the network into their routing table
* Routing table consists of buckets
 - Follow similar design as Kademlia k-buckets
 - Buckets have different max-sizes, depending on their usage
* Objects are the key-value pairs that nodes distribute

#### Routing (transport, RPC)
* Transport security is not incorporated into the design - it's ensured by CJDNS
* NAT is non-existent due to CJDNS
* Nodes communicate via QRP (simple RPC protocol) on top of SCTP (reliable, packet-based protocol)

### Content mechanisms
#### Replication Model
* TODO - push-pull or subscribe-update model
* Objects are replicated according to priority (a field determined by replicator which states the wanted propogation speed) and individual node assessment. 
#### Persistence
* Object persistence is determined locally. By default the lifetime of all objects is infinite. For update based applications (such as instant messaging), a shorter object expirary can be set locally when instantiating XDen
#### Load balancing
* TODO

### Node mechanisms
#### Reputation
* TODO
#### Flooding
* TODO

### API
* TODO