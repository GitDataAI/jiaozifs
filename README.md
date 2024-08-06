
# JiaoZiFS (JZFS)
A version control file system for data linage & data collaboration.

<p align="left">
  <a href="https://codecov.io/gh/jiaozifs/jiaozifs"><img src="https://codecov.io/gh/gitdataai/jiaozifs/branch/main/graph/badge.svg" /></a>
  <a href="https://goreportcard.com/report/github.com/jiaozifs/jiaozifs"><img src="https://goreportcard.com/badge/github.com/gitdataai/jiaozifs" /></a>  
  <a href=""><img src="https://img.shields.io/badge/golang-%3E%3D1.22.0-blue.svg" /></a>
  <br/>
</p>

<a href="https://github.com/GitDataAI/jiaozifs"><img src="https://github.com/GitDataAI/jiaozifs/blob/main/docs/logo/jiaozifs.png?raw=true" width="100" /></a>

----
### What is JiaoZiFS?
JiaoZiFS is an industry-leading **Data-Centric Version Control** File System, helps ensure Responsible AI Engineering by improving **Data Versioning**, **Provenance**, and **Reproducibility**.

Note:
* The name JiaoZi pays tribute to the world's earliest paper money: [Song Dynasty JiaoZi](https://en.wikipedia.org/wiki/Jiaozi_(currency)).
* JiaoZiFS is yet another implementation of [IPFS (InterPlanetary File System)](https://ipfs.tech/) as JiaoZiFS will be compatible with the [implementation requirements](https://specs.ipfs.tech/architecture/principles/#ipfs-implementation-requirements) of IPFS.
* As a filesystem of data versioning at scale, although JiaoZiFS is built for machine learning, It has a wide range of use scenarios (refer A Universe of Uses) and can be seamlessly integrated into all your data stack.

Data-centric AI is about the practice of iterating and collaborating on data, used to build AI systems, programmatically. Machine learning pioneer Andrew Ng [argues that focusing on the quality of data fueling AI systems will help unlock its full power](https://youtu.be/TU6u_T-s68Y).

----
### Why JiaoZiFS?
In production systems with machine learning components, updates and experiments are frequent. New updates to models(data products) may be released every day or every few minutes, and different users may see the results of different models as part of A/B experiments or canary releases.

* **Version Everything**: Data scientists are often criticized for being less disciplined with versioning their experiments(versioning of data, pipeline, code, and models), especially when using computational notebooks.
* **Track Data Provenance**: This applies to all processing steps in an AI/ML pipeline, including data collection/acquisition, data merging, data cleaning, feature extraction, learning, or deployment.
* **Reproducibility**: A final question of AI/ML that is often relevant for debugging, audits, and also science more broadly is to what degree data, models, and decisions can be reproduced.

----
### A Universe of Uses
JiaoZiFS's versatility shines across different industries – making it the multi-purpose tool for the **data centric applications and teams**.


* **Defining artificial intelligence in the context of lineage**: Artificial intelligence (AI) is an umbrella term that covers a variety of techniques and approaches that make it possible for machines to learn, adjust and act with intelligence comparable to the natural intelligence of humans. Lineage has direct implications for many of the techniques and approaches of AI, such as:
  * Neural networks. AI classifies data to make predictions and decisions in much the same way a human brain does. A neural network is a computing system made up of interconnected units (like neurons) that process data from external inputs, relaying information between each unit. The neural network requires multiple passes at the data to find connections and derive meaning from undefined data. Neural networks benefit greatly from the movement aspects of data lineage – because connecting those dots directs its search for meaning.
  * Natural language processing. AI that enables interaction, understanding and communication between humans and machines by analyzing and generating human language, including speech, is called natural language processing (NLP). NLP allows humans to communicate with computers using normal, everyday language to perform tasks. Natural language processing relies heavily on the human language data descriptions provided by the characteristics aspect of data lineage.
  * Machine learning. AI that’s focused on giving machines access to data and letting them learn for themselves is known as machine learning. Machine learning automates analytical model building using methods from neural networks, statistics, operations research and physics – and it finds hidden insights in data without being explicitly programmed where to look or what to conclude. Machine learning delves into the relationships, processes and transformations aspects of data lineage during its undirected exploration of data’s potential.
  * Deep learning. With deep learning, AI uses huge neural networks with many layers of processing to learn complex patterns in large amounts of data and perform humanlike tasks, such as recognizing speech or understanding images and videos (also known as computer vision). This method takes advantage of advances in computing power and improved training techniques. Deep learning depends on the users’ aspect of data lineage because its education is guided by analyzing how users interact with data.
* **Solving the Mysteries of Data Science’s Past and Present with Data Lineage**: Data lineage is an essential aspect of data science and data analytics that enables organizations to understand the journey of data from its origin to its destination. It is a process of tracking the origin, movement, and transformation of data through various stages of its lifecycle. Data lineage plays a critical role in enhancing the performance and productivity of data science and analytics teams.
  * By establishing a clear data lineage, data scientists and analysts can easily identify the source of data, its quality, and its dependencies. This information is crucial in ensuring data accuracy, reliability, and consistency. Additionally, data lineage helps teams quickly identify errors or issues in the data processing pipeline, enabling them to take corrective action promptly.
  * Moreover, data lineage promotes better collaboration among team members by providing a shared understanding of the data ecosystem. This shared understanding ensures that everyone involved in the data analysis process is on the same page, reducing confusion and errors caused by miscommunication.
* **Enterprise DataHub & Data Collaboration**: Depending on your operating scale, you may even be managing multiple team members, who may be spread across different locations. JiaoZiFS enable Collaborative Datasets Version Management at Scale,Share & collaborate easily: Instantly share insights and co-edit with your team.
* **DataOps & Data Products & Data Mesh**: Augmenting Enterprise Data Development and Operations,JiaoZiFS ensures Responsible DataOps/AIOps/MLOps by improving Data Versioning, Provenance, and Reproducibility. JiaoziFS makes a fusion of data science and product development and allows data to be containerized into shareable, tradeable, and trackable assets(data products or data NFTs). Versioning data products in a maturing Data Mesh environment via standard processes, data consumers can be informed about both breaking and non-breaking changes in a data product, as well as retirement of data products.
* **Industrial Digital Twin**: Developing digital twins for manufacturing involves managing tons of large files and multiple iterations of a project. All of the data collected and created in the digital twin process (and there is a lot of it) needs to be managed carefully. JiaoziFS allows you to manage changes to files over time and store these modifications in a database.
* **Data Lake Management**: Data lakes are dynamic.   New files and new versions of ex- isting files enter the lake at the ingestion stage.   Additionally, extractors can evolve over time and generate new versions of raw data.   As a result, data lake versioning is a cross-cutting concern across all stages of a data lake.   Of course vanilla distributed file systems are not adequate for versioning-related operations.   For example, simply storing all versions may be too costly for large datasets, and without a good version manager, just using filenames to track versions can be error-prone.   In a data lake, for which there are usually many users, it is even more important to clearly maintain correct versions being used and evolving across different users.   Furthermore, as the number of versions increases, efficiently and cost-effectively providing storage and retrieval of versions is going to be an important feature of a successful data lake system.


----
### Specification

[JiaoZiFS Specification](https://github.com/GitDataAI/Specification/blob/main/JiaoziFS)

----
### Basic Build And Usage

#### Requirement

1. To build JiaoZiFS, you need a working installation of   [Go 1.22.0 or higher](https://golang.org/dl/)
2. JiaoZiFS use postgres to store running data, you can install at  [postgres install installation guide](https://www.postgresql.org/docs/current/installation.html)

#### Build And Running

1. clone and build
```bash
git clone https://github.com/GitDataAI/jiaozifs.git
cd jiaozifs
make build
```

After following the above steps, you should be able to see an executable file named "jzfs."

2. init program and running
```bash
./jzfs init  --db postgres://<username>:<password>@localhost:5432/jiaozifs?sslmode=disable
./jzfs daemon
```

#### run with docker

```bash
docker run -v <data>:/app -p 34913:34913 gitdatateam/jzfs:latest  --db "postgres://<user>:<password>@192.168.1.16:5432/jiaozifs?sslmode=disable" --bs_path /app/data --listen http://0.0.0.0:34913 --config /app/config.toml
```

----
### Cloud

[Try without installing](https://cloud.jiaozifs.com)

Note: storage config for IPFS backend storage as you create a new repository in JiaoZiFS UI.

```
 {"type":"ipfs","ipfs":{"url":"/dns/kubo-service.ipfs.svc.cluster.local/tcp/5001"}}
```

### Build AL/ML pipeline over JiaoZiFS

[Face detection and recognition inference pipeline](https://colab.research.google.com/drive/1wsv-KMxTdsCLZ64eLq4W1MTfspid-vv6?usp=sharing)

----
### Contributors

<a href="https://github.com/hunjixin" target="_blank"><img src="https://avatars.githubusercontent.com/u/41407352?v=4" width="5%" height="5%"/> </a>
<a href="https://github.com/Brownjy" target="_blank"><img src="https://avatars.githubusercontent.com/u/54040689?v=4" width="5%" height="5%"/> </a>
<a href="https://github.com/TsumikiQAQ" target="_blank"><img src="https://avatars.githubusercontent.com/u/116857998?v=4" width="5%" height="5%"/> </a>
<a href="https://github.com/taoshengshi" target="_blank"><img src="https://avatars.githubusercontent.com/u/33315004?v=4" width="5%" height="5%"/> </a>
<a href="https://github.com/gitdata001" target="_blank"><img src="https://avatars.githubusercontent.com/u/157772574?v=4" width="5%" height="5%"/> </a>

----
### Our Users
JiaoZiFS is supported or integrated by the following companies/projects. If you want to be listed here, please contact us.

* DeSci Asia
* 6079.AI
* MesoReefDAO

### Our Partners
JiaoZiFS is supported or integrated by the following companies/projects. If you want to be listed here, please contact us.  
* Artizen
* Functionland
* DeSci Asia
  
----
### License

Dual-licensed under [MIT](https://github.com/GitDataAI/jiaozifs/blob/main/LICENSE-MIT) + [Apache 2.0](https://github.com/GitDataAI/jiaozifs/blob/main/LICENSE-APACHE)



Fri Jul 26 01:06:29 UTC 2024
Mon Jul 29 01:09:05 UTC 2024
Tue Jul 30 01:07:19 UTC 2024
Wed Jul 31 00:56:12 UTC 2024
Thu Aug  1 01:13:17 UTC 2024
Fri Aug  2 01:07:13 UTC 2024
Mon Aug  5 01:09:30 UTC 2024
Tue Aug  6 01:07:40 UTC 2024
