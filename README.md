# XDen
XDen is a P2P data store that offers 256bit keys associated with mutable data. It is simple and efficient. It is designed to supercede common faults of a P2P system through mechanisms such as reputation management and modularization of security. It is developed in the Go programming language, a compiled and concise language that offers high-level language features such as garbage collection and extension libraries (JSON parsing etc).  
XDen will be the first application designed to run over CJDNS. CJDNS is a routing engine designed for speed, security and scalability. It guarantees end-to-end security of data, no interefering NAT and that everyone has an IPv6 address. 

## Get involved
To become involved, just go to the issues page of this repository. Development is done in the /development branch, which is frequently updated. When XDen reaches a somewhat stable compilation level (that is, tests can run) data will be pushed to /master.

Make sure to read up on the [design doc](https://github.com/liamzebedee/xden/wiki/Design). If you want some extra info, just contact me, liamzebedee.

# Notes
P2P is a movement concerning the decentralization of common Internet technologies (WWW, IRC, Filesharing) through mesh networking. I have become widely interested and involved in such movements over the past year, researching and developing ideas surrounding P2P technologies such as Tor, Freenet, I2P and DHTs in general. While such technologies are ok, much can be done to improve them:
*  Security: Almost every P2P app has implemented an independent transport security layer, which to some isn’t fully necessary, if a common networking layer is implemented. 
*	Security: The basis for most decentralized security systems is public key cryptography. Being a fundamental part of P2P, still only a few apps use Elliptic Curve Cryptography, a more efficient algorithm than RSA that provides smaller key sizes for equivalent protection
*	Configurability: Optimizing operations for different environments improves efficiency (and being a mesh network, this improves the entire networks performance). 
*	Documentation: If there is ever going to be any forward movement, it is through learning, which is hard if you don’t create documentation. 
*	Simplicity