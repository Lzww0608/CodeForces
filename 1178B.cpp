#include <vector>
#include <iostream>
#include <algorithm>
#include <stack>
#include <cmath>
#include <string>
using namespace std;


int main() {
	string s;
	cin >> s;
	long long ans = 0, n = s.length();
	long long pre = 0, suf = 0;
	for (int i = 1; i < n - 1; ++i) {
		if (s[i] == 'v' && s[i + 1] == 'v') {
			suf++;
		}
	}
	for (int i = 1; i < n - 2; ++i) {
		if (s[i] == 'o') {
			ans += pre * suf;
		} else {
			if (s[i-1] == 'v') 
				pre++;
			if (s[i+1] == 'v')
				suf--;
		}
	}
	cout << ans << endl;
	return 0;
}

