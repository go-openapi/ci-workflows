:scan_input {
    # main scanner loop, looking for "Misspelled words:" entries
    /^Misspelled words:/,/^-\{5,\}/ {
        /^&lt;html\(attribute\|content\)&gt\;/ {
            # Extract the file name
            # TODO: be less strict here
            s/^&lt;html\(attribute\|content\)&gt\;\s\+\(.\+\.md\):.*$/"\2": [/
            h
            n  # skip current match
            n  # skip next line
    
            # go decode a section of single words, enclosed between "=====..." lines
            b parse_section
        }
    
        n
    
        b scan_input
    }
}

# Back to main loop (cycles)
d

:parse_section {
    :word /^[^-]/ {
        # Append words to the previous line
        # Trim space around words
        s/\s\+//g
        # Quote word
        s/^/"/
        s/$/"/

        # Add comma separator
        s/$/, /
        H
        n

        b word
    }

    /^-\{5,\}$/ {
        # Print the collected words as a comma-separated array of words
        x

        s/^/{/g
        s/\n/ /g
        s/\s\+/ /g
        s/\s\+$//g
        s/\(,\s*\)\?$/ ]/
        s/$/}/g

        p

        # Start a new cycle
        d
    }
}

d
