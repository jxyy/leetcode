#include<iostream>
#include<vector>
using namespace std;

class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* next;

    Node() {}

    Node(int _val, Node* _left, Node* _right, Node* _next) {
        val = _val;
        left = _left;
        right = _right;
        next = _next;
    }
};

class Solution {
public:
    Node* connect(Node* root) {
        if(root == NULL){
            return NULL;
        }
        auto level = vector<Node*>();
        level.push_back(root);
        while (level.size() > 0)
        {
            auto next_level = vector<Node*>();
            for(int i = 0; i < level.size(); i ++){
                auto node = level[i];
                if(node->left != NULL) {
                    next_level.push_back(node->left);
                    next_level.push_back(node->right);
                }
                if(i < level.size()-1) {
                    node->next = level[i+1];
                }
            }
            level = next_level;
        }
        return root;
    }
};

int main(){
    return 0;
}
