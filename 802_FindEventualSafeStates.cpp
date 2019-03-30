#include <iostream>
#include <vector>
#include <set>
#include <queue>
using namespace std;

class Solution {
public:
    //solution two: BFS, 类似拓扑排序
    vector<int> eventualSafeNodes(vector<vector<int>>& graph) {
        int size = graph.size();
        vector<set<int> > newGraph(size, set<int>());
        vector<set<int> > reverseGraph(size, set<int>());
        queue<int> q;
        vector<bool> safe(size, false);
        vector<int> res;
        
        for (int i = 0; i < size; i++) {
            if (graph[i].size() == 0) {
                q.push(i);
            }
            for (auto n : graph[i]) {
                newGraph[i].insert(n);
                reverseGraph[n].insert(i);
            }
        }
        
        while(!q.empty()) {
            int top = q.front();
            q.pop();
            safe[top] = true;
            for (auto n : reverseGraph[top]) {
                newGraph[n].erase(top);
                if (newGraph[n].empty()) {
                    q.push(n);
                }
            }   
        }
        for (int i = 0; i < size; i++) {
            if (safe[i]) {
                res.push_back(i);
            }
        }
        return res;
    }
};
