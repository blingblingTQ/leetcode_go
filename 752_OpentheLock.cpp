#include <iostream>
#include <queue>
#include <unordered_map>
using namespace std;

class Solution {
public:
    vector<string> nextStep(string str, unordered_map<string, bool>& seen) {
        vector<string> res;
        vector<int> mod = {1, -1};
        for(int k = 0; k < 2; k++) {
            for(int i = 0; i < 4; i++) {
                string nextStr = str;
                char nextChar = nextStr[i]+mod[k];
                if(nextChar == '0'+10) {
                    nextChar = '0';
                }
                nextStr[i] = nextChar;
                if(!seen[nextStr]) {
                    res.push_back(nextStr);
                }
            }
        }
        return res;
    }
    
    int openLock(vector<string>& deadends, string target) {
        string initStr = "0000";
        queue<string> q;
        int step = 0;
        unordered_map<string, bool> seen;
        for(string str : deadends) {
            seen[str] = true;
        }
        
        q.push(initStr);
        while(!q.empty()) {
            int size = q.size();
            for(int i = 0; i < size; i++) {
                string str = q.front();
                q.pop();
                if(seen[str]) {
                    continue;
                }
                if(str == target) {
                    return step;
                }
                
                seen[str] = true;
                vector<string> nextStrs = nextStep(str, seen);
                for(auto s : nextStrs) {
                    if(seen[s]) {
                        continue;
                    }
                    q.push(s);
                }
            }
            step++;
        }
        return -1;
    }
};



int main(void) {
    Solution solution;
    vector<string> deadends = {"0201", "0101", "0102", "1212", "2002"};
    string target = "0202";
    cout << solution.openLock(deadends, target) << endl;
    return 0;
}

