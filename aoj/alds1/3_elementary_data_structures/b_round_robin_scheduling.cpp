#include <iostream>
#include <queue>

using namespace std;

typedef struct {
  string name;
  int time;
} Process;

int main() {
  int n, q;
  queue<Process> que;
  
  cin >> n >> q;
  for (int i = 0; i < n; i++) {
    string name;
    int time;

    cin >> name >> time;

    Process p = {name, time};
    que.push(p);
  }

  int sum = 0;
  while (!que.empty()) {
    Process pro = que.front();
    que.pop();
    
    if (pro.time <= q) {
      sum += pro.time;
      cout << pro.name << " " << sum << endl;
    } else {
      sum += q;
      pro.time -= q;
      que.push(pro);
    }
  }
  
  return 0;
}
