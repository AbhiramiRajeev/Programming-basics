#!/bin/bash

# Define base directory
BASE_DIR=~/Documents/Programming

# Define problem names correctly in the right order
problems=(
    "1_remove_occurrences" "2_max_element" "3_reverse_array" "4_remove_duplicates"
    "5_count_occurrences" "6_max_element_again" "7_reverse_array_again" "8_is_sorted"
    "9_first_unique_element" "10_max_difference" "11_first_unique_char" "12_array_intersection"
    "13_majority_element" "14_first_unique_char_again" "15_remove_duplicates_sorted"
    "16_move_zeros" "17_missing_number" "18_find_duplicates" "19_longest_consecutive"
    "20_subarray_sum"
)

# Define folders and file extensions
folders=("Golang" "JavaScript" "Python")
extensions=(".go" ".js" ".py")

# Loop through folders and create files
for i in "${!folders[@]}"; do
    folder="${folders[$i]}"
    ext="${extensions[$i]}"
    
    # Ensure the folder exists
    mkdir -p "$BASE_DIR/$folder"

    # Create files with question number and problem name
    for problem in "${problems[@]}"; do
        touch "$BASE_DIR/$folder/$problem$ext"
    done
done

echo "Files created successfully!"

# Additional logic for Golang folder: Create subfolders and move files
GOLANG_DIR="$BASE_DIR/Golang"

for problem in "${problems[@]}"; do
    # Extract short name after the number (removes "1_", "2_", etc.)
    short_name=$(echo "$problem" | cut -d'_' -f2-)
    
    # Create subfolder inside Golang folder
    mkdir -p "$GOLANG_DIR/$short_name"

    # Move the corresponding .go file into the newly created subfolder
    mv "$GOLANG_DIR/$problem.go" "$GOLANG_DIR/$short_name/"
done

echo "Golang folder structured successfully!"

