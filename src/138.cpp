#include<iostream>
#include<vector>
#include<unordered_map>
using namespace std;


// Definition for a Node.
class Node {
public:
    int val;
    Node* next;
    Node* random;

    Node() {}

    Node(int _val, Node* _next, Node* _random) {
        val = _val;
        next = _next;
        random = _random;
    }
};


class Solution {
public:
    Node* copyRandomList(Node* head) {
        if(head==NULL) return NULL;

        unordered_map<Node*, Node*> map;
        Node* new_head = new Node(head->val, NULL, NULL);
        map[head] = new_head;
        Node* node = head, *new_node = new_head;
        while (node->next!=NULL){
            Node* new_next = new Node(node->next->val, NULL, NULL);
            new_node->next = new_next;
            map[node->next] = new_next;
            node = node->next;
            new_node = new_next;
        }

        node = head;
        while (node!=NULL){
            if(node->random != NULL) {
                map[node]->random = map[node->random]; 
            }
            node = node->next;
        }

        return new_head;
    }
};

int main(){
    return 0;
}
