#include <iostream>
#include <vector>
#include <queue>
#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);
  
  int n;
  cin >> n;
  
  vector<vector<int> > tree(n+1, vector<int>());

  int a, b;
  for (int i = 0; i < n-1; i++) {
    cin >> a >> b;
    tree[a].push_back(b);
    tree[b].push_back(a);
  }
  
  priority_queue<int, vector<int>, greater<int> > que;
  que.push(1);

  vector<bool> visit(n+1, false);
  while (que.size()) {
    int i = que.top();
    que.pop();

    visit[i] = true;
    
    vector<int> links = tree[i];
    for (int j = 0; j < links.size(); j++) {
      int k = links[j];
      if (!visit[k]) {
        que.push(k);
      }
    }

    if (i != 1) {
      cout << " ";
    }
    cout << i;
  }
  cout << eol;
  
  return 0;
}
