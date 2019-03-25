class Employee {
public:
    // It's the unique ID of each node.
    // unique id of this employee
    int id;
    // the importance value of this employee
    int importance;
    // the id of direct subordinates
    vector<int> subordinates;
};

class Solution {
public:

    int dfs(int id, unordered_map<int, Employee*> &map) {
        Employee *em = map[id];
        int sum = em->importance;
        for(int i = 0; i < em->subordinates.size(); i++) {
           sum += dfs(em->subordinates[i], map);
        }
        return sum;
    }
    
    int getImportance(vector<Employee*> employees, int id) {
        std::unordered_map<int, Employee*> map;
        for(auto employee : employees) {
            map[employee->id] = employee;
        }
        return dfs(id, map);
    }
};
