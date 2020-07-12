package prime 

import (
  "github.com/tmthrgd/go-memset"
  "fmt"
)

const MAX = 10000

int prefix[MAX + 1]; 
  
func BuildPrefix() { 
    // Create a boolean array "prime[0..n]". A  
    // value in prime[i] will finally be false  
		// if i is Not a prime, else true. 
		// s := make([]boolean, MAX)
    fmt.Println("Inside buildPrefix ...")
    // bool prime[MAX + 1]; 
    // memset(prime, true, sizeof(prime)); 
  
    // for (int p = 2; p * p <= MAX; p++) { 
  
    //     // If prime[p] is not changed, then  
    //     // it is a prime 
    //     if (prime[p] == true) { 
  
    //         // Update all multiples of p 
    //         for (int i = p * 2; i <= MAX; i += p) 
    //             prime[i] = false; 
    //     } 
    // } 
  
    // // Build prefix array 
    // prefix[0] = prefix[1] = 0; 
    // for (int p = 2; p <= MAX; p++) { 
    //     prefix[p] = prefix[p - 1]; 
    //     if (prime[p]) 
    //         prefix[p]++; 
    // } 
} 

