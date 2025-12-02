#include <iostream>
#include <string>

// Regular single-line comment
/// Doxygen documentation comment
//! Important Doxygen documentation
///< Doxygen end-of-line comment
//!< Another Doxygen end-of-line style

/**
 * Main Doxygen documentation block
 * @brief Example program
 * @param argc Number of arguments
 * @param argv Array of arguments
 * @return Exit code
 */
 
/*!
 * Alternative Doxygen style
 * Important documentation
 */

/* Regular multi-line
   comment that should
   be removed completely */

class Example {
    //! Doxygen property documentation
    int value;
    
    /// Doxygen method documentation
    void test() {
        // Regular comment inside method
        std::string s = "/* This is not a comment */";
        char c = '//';  /* Also not a comment */
        std::cout << "Hello /* world */" << std::endl; // End of line comment
    }
};

/* Multi-line comment with nested symbols
   // Like this one
   /* Or this */


int main() {
    /// @brief Main function doc
    return 0;  //! Important note
}