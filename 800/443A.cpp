#include <iostream>
#include <vector>
#include <unordered_map>
#include <cstdlib>
#include <cmath>
#include <algorithm>
#include <string>
using namespace std;
using ll = long long;


int main() {
	std::ios::sync_with_stdio(false);
	cin.tie(nullptr);
	cout.tie(nullptr);

	string s;
	getline(cin, s);
	unordered_map<char, bool> m;
	int ans = 0;
	for (int i = 1; i < s.length() - 1; i += 3) {
		if (m[s[i]] == false) {
			ans++;
			m[s[i]] = true;
		}
	}

	cout << ans << endl;

	return 0;
}
