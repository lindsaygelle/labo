# Nintendo Labo


Labo is a fan project that fetches and wrangles information about the Nintendo Labo product suite. Package draws upon the Nintendo official store and Labo official website to build out the package struct contents. Functions are designed to try collect the contextual information for each Nintendo Labo product and express each edge verbosity. Can be used to build out a basic HTTP API or CLI that can let users lookup Labo information (for example).

Package offers Nintendo Labo information at two levels of detail. A `labo.Labo` struct contains the common fields for a Nintendo Labo product that is sold on the Nintendo store. The second is the `labo.Kit` which contains the unique information found on the official website. 

Package also offers some basic utils for wrangling the struct data, such as marshalers and unmarshalers. Any contributions are welcome. Is stable for as long as the Nintendo store uses the same page layout.


! Make more Labo's Nintendo !


#### Nintendo scrape resources
###### Nintendo store
[https://store.nintendo.com]
###### Labo official website
[https://labo.nintendo.com]

--
[🗾][🗃️]
