#include <iostream>
#include <string>


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





class Example {
    //! Doxygen property documentation
    int value;
    
    /// Doxygen method documentation
    void test() {
        
        std::string s = "/* This is not a comment */";
        char c = '//';  
        std::cout << "Hello /* world */" << std::endl; 
    }
};






int main() {
    /// @brief Main function doc
    return 0;  //! Important note
}