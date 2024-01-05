

struct Node{
int value;
Node* next;
};

Node list_insert(*Node node, int n)
{
    struct Node newNode;
    newNode.value = n;
    node.next = &newNode;

    return newNode
}

void main()
{
    struct Node initNode;
    struct Node node;

        node = list_insert(initNode, 1);
        node = list_insert(node, 2);
        node = list_insert(node, 3);
        node = list_insert(node, 4);
        node = list_insert(node, 5);

    return
}