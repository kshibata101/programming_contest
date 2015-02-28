#include <iostream>
#include <vector>
#define eol '\n';
using namespace std;

struct UnionFind {
  vector<int> data;
  UnionFind(int size): data(size, -1) {}
  bool unionSet(int x, int y) {
    x = root(x); y = root(y);
    if (x != y) {
      if (data[y] < data[x]) swap(x, y);
      data[x] += data[y];
      data[y] = x;
    }
    return x != y;
  }
  bool findSet(int x, int y) {
    return root(x) == root(y);
  }
  int root(int x) {
    return data[x] < 0 ? x : root(data[x]);
  }
  int size(int x) {
    return -data[root(x)];
  }
  bool isRoot(int x) {
    return data[x] < 0;
  }
};


int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int n;
  int num = 26;
  UnionFind uft(num);
  int* edges = new int[num];

  cin >> n;

  string line;
  for (int i = 0; i < n; i++) {
    cin >> line;
    
    char s = line[0];
    char e = line[line.size() - 1];
    int is = int(s - 'a');
    int ie = int(e - 'a');
    edges[is]++;
    edges[ie]++;
    
    uft.unionSet(is, ie);
  }
  
  // count
  int ans = 0;

  // まず木の数を数える
  vector<int> tree_root;
  for (int i = 0; i < num; i++) {
    if (uft.isRoot(i) && edges[i] > 0) {
      tree_root.push_back(i);
    }
  }
  ans += tree_root.size() - 1;

  // 次に木ごとに一筆書きができるか考える
  for (int j = 0; j < tree_root.size(); j++) {
    int root = tree_root[j];
    int odd_edge_num = 0;
    for (int i = 0; i < num; i++) {
      if (uft.root(i) == root && edges[i] % 2 != 0) {
        odd_edge_num++;
      }
    }
    ans += max(odd_edge_num / 2 - 1, 0);
  }

  cout << ans << eol;

  return 0;
}
