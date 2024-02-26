#include <vector>
#include <iostream>
#include <algorithm>
#include <stack>
#include <cmath>
#include <string>
using namespace std;

int main() {
	std::ios::sync_with_stdio(false);
	cin.tie(nullptr);
	int t;
	cin >> t;
	string s;
	while (t--) {
		cin >> s;
		stack<char> st;
		int n = s.length();
		string ans;
		for (int i = 0; i < n; ++i) {
			if (st.empty() || st.top() <= s[i]) {
				st.push(s[i]);
			}
			else {
				while (!st.empty() && st.top() > s[i]) {
					char c = st.top();
					int t = c - '0' + 1;
					st.pop();
					ans.push_back(char(min(t, 9) + '0'));
				}
				st.push(s[i]);
			}
		}
		while (!st.empty()) {
			ans += st.top();
			st.pop();
		}
		sort(ans.begin(), ans.end());
		cout << ans << endl;
	}
	return 0;
}