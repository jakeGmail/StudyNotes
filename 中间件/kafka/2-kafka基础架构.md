[toc]

# 1 kafka基础架构
1.为方便扩展，并提高吞吐量,一个topic分为多个partition
⒉.配合分区的设计，提出消费者组的概念，组内每个消费者并行消费。一个分区的数据只能由同组的一个消费者来消费。
3.为提高可用性，为每个partition增加若干副本，是一种Master-Slave模式，Slave是用于备份Master的数据。当Master挂了，就由Slave在充当master的角色。
4.ZK(Zookeeper，等同于rokerMQ的NameServer)中记录全部leader信息，还会帮我们存储哪一个分区下谁是leader。Kafka2.8.0以后也可以配置不采用ZK，即Zookeeper是可选的