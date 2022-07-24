# Lush Climate Change Data Server / Client

**Decentralized storage and file management for climate data**

The Paris Climate Accord was a momentous accomplishment, uniting all the world's nations in an effort to tackle climate change. ClimateDataPool intends to unite the reporting of all the nations on a blockchain.

Our objective with this project is to build an IPFS/Filecoin-based file management app for collecting and sharing climate data, and then go a step further to help organize and visualize the data for the end consuner, which can be Green Tech companies or Web3 DAO's

The technical flow of this diagram is as follows

1. The go-server runs a GRPC/API client, which will listen to and update the IPFS data store with the most up-to-date information, available from various sources like IoT Endpoints or National Electricity Grid
2. This data will be served over our front-end to companies / DAO's / citizens for consumption
3. We are also buliding a CLI application in Go, which can be used by other servers around the world to consume the latest information about Climate Change, thereby building their own use-cases

We are on democratizing the Climate Change Data via the Decentralized web, so as to empower other developers and companies to build greener tech

-- Insert Image Here --
