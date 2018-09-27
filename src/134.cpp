#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:
    int canCompleteCircuit(vector<int>& gas, vector<int>& cost) {
        int n = gas.size();
        int count = 0, pass_count = 0;
        int start = 0, end = 0;
        int gas_till_now = 0, cost_till_now = 0;
        while(count < n && pass_count < n) {
            while(count < n && gas_till_now + gas[end] >= cost_till_now + cost[end]) {
                count ++;
                pass_count ++;
                gas_till_now += gas[end];
                cost_till_now += cost[end];
                end = (end+1)%n;
            }
            if(count < n){
                end = (end+1)%n;
                start = end;
                count = gas_till_now = cost_till_now = 0;
                pass_count ++;
            }
        }
        return count == n ? start : -1;
    }
};

int main(){
    Solution s;
    int raw_gas[] = {1,2,3,4,5};
    int raw_cost[] = {3,4,5,1,2};
    vector<int> gas(raw_gas, raw_gas+5);
    vector<int> cost(raw_cost, raw_cost+5);
    cout << s.canCompleteCircuit(gas, cost) << endl;

    int raw_gas2[] = {3,3,4};
    int raw_cost2[] = {3,4,4};
    vector<int> gas2(raw_gas2, raw_gas2+3);
    vector<int> cost2(raw_cost2, raw_cost2+3);
    cout << s.canCompleteCircuit(gas2, cost2) << endl;

    int raw_gas3[] = {7,1,0,11,4};
    int raw_cost3[] = {5,9,1,2,5};
    vector<int> gas3(raw_gas3, raw_gas3+5);
    vector<int> cost3(raw_cost3, raw_cost3+5);
    cout << s.canCompleteCircuit(gas3, cost3) << endl;
    return 0;
}
