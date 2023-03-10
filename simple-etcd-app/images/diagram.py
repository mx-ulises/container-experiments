from diagrams import Cluster, Diagram, Edge
from diagrams.onprem.client import Users
from diagrams.onprem.container import Docker
from diagrams.onprem.network import Etcd

graph_attr = {
    'bgcolor': 'transparent'
}

with Diagram('Simple ETCD App', show=False, graph_attr=graph_attr):
    user = Users('Users')

    with Cluster('Backend Services'):
        my_app = Docker('MyApp')
        etcd = Etcd('etcd')

        user >> Edge(label='Read') >> my_app
        user >> Edge(label='Increase') >> my_app
        my_app >> etcd
