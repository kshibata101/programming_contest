#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

// input 
int n = 4;
vector<int> a = {1,2,4,7};
int k = 13; // => Yes
// int k = 15 // => NO

bool dfs(int sum, int index)
{
  if (index == n) {
    return (sum == k);
  }

  if (dfs(sum, index+1)) return true;
  if (dfs(sum + a[index], index+1)) return true;
  
  return false;
}

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  // main
  if (dfs(0,0)) {
    cout << "Yes" << eol;
  } else {
    cout << "No" << eol;
  }
  
  return 0;
}


