#include <iostream>
#include <vector>
#include <algorithm>
#define eol '\n'
using namespace std;

vector<int> perm;
int n;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  cin >> n;

  perm.resize(n);
  for (int i = 0; i < n; i++) {
    perm[i] = i;
  }
  do {
    cout << perm[0];
    for (int i = 1; i < n; i++) {
      cout << " " << perm[i];
    }
    cout << eol;
  } while (next_permutation(perm.begin(), perm.end()));
  return 0;
}
