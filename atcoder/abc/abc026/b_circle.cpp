#include <iostream>
#include <vector>
#include <algorithm>
#include <math.h>
#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int N;
  cin >> N;

  vector<int> R;
  for (int i = 0; i < N; i++) {
    int tmp;
    cin >> tmp;
    R.push_back(tmp);
  }

  sort(R.begin(), R.end(), greater<int>());

  int sum = 0;
  for (int i = 0; i < N; i++) {
    if (i % 2 == 0) {
      sum += R[i] * R[i];
    } else {
      sum -= R[i] * R[i];
    }
  }
  printf("%.7lf\n", sum * M_PI);
  
  return 0;
}
