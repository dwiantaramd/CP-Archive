#include <bits/stdc++.h>

using namespace std;

int main() {
    long n;
    cin >> n;

    long x[n*4];
    for(long i = 0; i<n*4; i++) {
        cin >> x[i];
    }
    int result = 0;
    for(long i = 0; i<n-1; i++) {
        long lebar = abs(x[(i*4)+2] - x[i*4]);
        long panjang = abs(x[(i*4)+3] - x[(i*4)+1]);
        for(long j = i+1; j<n;j++) {
            long Clebar = abs(x[(j*4)+2] - x[j*4]);
            long Cpanjang = abs(x[(j*4)+3] - x[(j*4)+1]);
            if(Clebar!= 0 && Cpanjang !=0 && lebar!= 0 && panjang != 0 && !((Clebar <= lebar && Cpanjang <= panjang) || (Clebar <= panjang && Cpanjang <= lebar) || (lebar <= Clebar && panjang <= Cpanjang) || (panjang <= Clebar && lebar <= Cpanjang))) {
                result++;
            }
        }

    }
    cout << result;

}
