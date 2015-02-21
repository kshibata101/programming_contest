#include <iostream>

using namespace std;

int main() {
  int n;
  cin >> n;

  int** G = new int*[n+1];
  for (int i = 0; i < n; i++) {
    int u,k,v;
    cin >> u >> k;
    G[u] = new int[n+1];
    for (int j = 0; j < k; j++) {
      cin >> v;
      G[u][v] = 1; 
    }
  }

  for (int i = 1; i <= n; i++) {
    cout << G[i][1];
    for (int j = 2; j <= n; j++) {
      cout << " " << G[i][j];
    }
    cout << endl;
  }

  return 0;
}
