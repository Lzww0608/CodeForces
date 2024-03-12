#include <cstdio>
#include <cstring>
#include <vector>
#include <iostream>
#include <cmath>
#include <algorithm>
using namespace std;
typedef long long ll;

int main(){
	std::ios::sync_with_stdio(false);
	cin.tie(nullptr); 
	cout.tie(nullptr);

	int n;
	cin >> n;
	vector<int> a(n), ans;
	int neg = INT_MIN, num = 0, zero = 0;
	for (int& x : a) {
		cin >> x;
		if (x < 0) {
			num++;
			neg = max(neg, x);
		}
		else if (x == 0) zero++;
		
	}
	if (n == 1) {
		cout << a[0] << endl;
		return 0;
	}
	
	if (zero == n || (zero == n - 1 && num == 1)) {
		cout << 0;
		return 0;
	}

	if (num & 1) {
		for (int x : a) {
			if (x == neg) {
				neg = 0;
			} else if (x != 0) {
				ans.push_back(x);
			}
		}
	} else {
		for (int x : a) {
			if (x != 0)
				ans.push_back(x);
		}
	}
	for (int x : ans)
		cout << x << " ";
	return 0;
	
}