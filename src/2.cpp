#include<iostream>
using namespace std;

struct ListNode{
    int val;
    ListNode *next;
    ListNode(int x): val(x), next(NULL){}
};

class Solution{
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2){
        ListNode* head = NULL, *tail = NULL;
        ListNode* num1 = l1, *num2 = l2;
        int addon = 0;
        while(num1 != NULL || num2 != NULL){
            int val1 = (num1!=NULL? num1->val : 0);
            int val2 = (num2!=NULL? num2->val : 0);
            int single_sum = val1 + val2 + addon;
            ListNode* new_num = new ListNode(single_sum%10);
            addon = single_sum/10;

            if(tail == NULL){
                head = tail = new_num;
            }else{
                tail->next = new_num;
                tail = new_num;
            }

            num1 = (num1!=NULL? num1->next : num1);
            num2 = (num2!=NULL? num2->next : num2);
        }

        if(addon != 0){
            ListNode* new_num = new ListNode(1);
            tail->next = new_num;
            tail = new_num;
        }
        return head;
    }
};

int main(){
    return 0;
}
