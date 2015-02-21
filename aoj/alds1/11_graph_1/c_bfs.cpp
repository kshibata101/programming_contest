#include <iostream>
#include <queue>

using namespace std;

typedef pair<int, int> E;

int INF = -1;

int main() {
  int n, u, k;
  cin >> n;
  int** G = new int*[n+1];
  for (int i = 0; i < n; i++) {
    cin >> u >> k;
    G[u] = new int[k+1];
    G[u][0] = k;
    for (int j = 1; j <= k; j++) {
      cin >> G[u][j];
    }
  }

  int cnt = 0;
  int* visit = new int[n+1];
  for (int i = 1; i <= n; i++) {
    visit[i] = INF;
  }

  queue<E> q;
  q.push(E(1, cnt));
  while (q.size()) {
    E e = q.front();
    q.pop();
    if (visit[e.first] != INF)
      continue;
    visit[e.first] = e.second;

    int len = G[e.first][0];
    for (int i = 1; i <= len; i++) {
      q.push(E(G[e.first][i], e.second + 1));
    }
  }

  for (int i = 1; i <= n; i++) {
    cout << i << " " << visit[i] << endl;
  }
  return 0;
}
