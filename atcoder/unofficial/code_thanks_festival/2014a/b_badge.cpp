#include <algorithm>
#include <vector>
#include <iostream>
#include <functional>
using namespace std;
     
int main() {
  int n, tmp;
  vector<int> list;
     
  cin >> n;
  for (int i = 0; i < 3; i++) {
    cin >> tmp;
    list.push_back(tmp);
  }
  sort(list.begin(), list.end(), greater<int>());
  int sum = list[0] + list[1] + list[2];
  int sho = n / sum;
  int amari = n % sum;
  int ans = sho * 3;
  if (amari == 0) {
  } else if (amari <= list[0]) {
    ans += 1;
  } else if (amari <= list[0] + list[1]) {
    ans += 2;
  } else {
    ans += 3;
  }
     
  cout << ans << endl;
  return 0;
}
