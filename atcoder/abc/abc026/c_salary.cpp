#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int get_salary(vector< vector<int> > map, int num) {
  if (map[num].size() == 0) {
    return 1;
  } else if (map[num].size() == 1) {
    return 2 * get_salary(map, map[num][0]) + 1;
  } else {
    int mx = -2147483648;
    int mn = 2147483647;
    
    for (int i = 0; i < map[num].size(); i++) {
      int sal = get_salary(map, map[num][i]);
      mx = max(sal, mx);
      mn = min(sal, mn);
    }
    return mx + mn + 1;
  }
}

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int N;
  cin >> N;

  vector< vector<int> > map(N+1);

  for (int i = 1; i < N; i++) {
    int tmp;
    cin >> tmp;
    map[tmp].push_back(i+1);
  }

  for (int i = 1; i < N; i++) {
    
  }

  cout << get_salary(map, 1) << eol;
  
  return 0;
}
