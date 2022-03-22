

 - Take in arbitrary input and return a deterministic color
 - Color chosen can be limited in several ways:
   - only visually / noticibly distinct colors to choose from
   - Color exclusions
     - dynamic color exclusions (optional terminal context)
   - colors within different terminal support classes (i.e. term-256)

 - Offer to return Hex codes (6 digits or 3)
 - Offer to return ascii escape codes
 - If the input is text, offer to wrap the input text and return the output as a string


1. take input as bytes
1. md5 hash the input
1. use modulo against the sum to choose the color to return from the subset of colors selected.
