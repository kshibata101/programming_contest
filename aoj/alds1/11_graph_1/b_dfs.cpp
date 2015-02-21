#include <iostream>

using namespace std;

typedef pair<int, int> p;

int n,u,k,t;
int** v;
p* T;

void bfs(int pos) {
  // if visited
  if (T[pos].first) {
    return;
  }
  // check visit
  T[pos].first = ++t;

  int leng = v[pos][0];
  for (int i = 1; i <= leng; i++) {
    bfs(v[pos][i]);
  }
  T[pos].second = ++t;
}

int main() {
  cin >> n;

  v = new int*[n+1];
  for (int i = 0; i < n; i++) {
    cin >> u >> k;
    v[u] = new int[k+1];
    v[u][0] = k;
    for (int j = 1; j <= k; j++) {
       cin >> v[u][j];
    }
  }

  T = new p[n+1];
  for (int i = 1; i <= n; i++) {
    T[i] = p(0, 0);
  }

  t = 0;
  for (int i = 1; i <= n; i++) {
    bfs(i);
  }

  for (int i = 1; i <= n; i++) {
    cout << i << " " << T[i].first << " " << T[i].second << endl;
  }
  
  return 0;
}
