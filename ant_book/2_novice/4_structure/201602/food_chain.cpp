#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

struct UnionFind {
  vector<int> data;
  UnionFind(int size): data(size, -1) {}
  bool unite(int x, int y) {
    x = root(x); y = root(y);
    if (x != y) {
      if (data[y] < data[x]) swap(x, y);
      data[x] += data[y];
      data[y] = x;
    }
    return x != y;
  }
  bool find(int x, int y) {
    return root(x) == root(y);
  }
  int root(int x) {
    return data[x] < 0 ? x : root(data[x]);
  }
  int size(int x) {
    return -data[root(x)];
  }
};

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int n,k;
  cin >> n >> k;
  
  UnionFind uf(3*n);
  int cnt = 0;
  
  for (int i = 0; i < k; i++) {
    int t,x,y;
    cin >> t >> x >> y;

    if (x <= 0 || x > n || y <= 0 || y > n) {
      cnt++;
      continue;
    }

    if (t == 1) {
      if (uf.find(x,y+n) || uf.find(x,y+2*n)) {
        cnt++;
        continue;
      }
      
      uf.unite(x,y);
      uf.unite(x+n,y+n);
      uf.unite(x+2*n,y+2*n);
    } else if (t == 2) {
      if (uf.find(x,y) || uf.find(x,y+2*n)) {
        cnt++;
        continue;
      }
      uf.unite(x,y+n);
      uf.unite(x+n,y+2*n);
      uf.unite(x+2*n,y);
    }
  }

  cout << cnt << eol;
  
  return 0;
}
