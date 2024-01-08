<?php
session_start();

// Check if the form is submitted
if ($_SERVER["REQUEST_METHOD"] == "POST") {
    // collect values of input fields
    $result = isset($_POST['name']) ? $_POST['name'] : null;

    // Check if $result is set before using it
    if (isset($result)) {
        // Retrieve the JSON data from the API URL
        $apiUrl = "http://localhost:8080/shop/users/" . urlencode($result);
        $jsonData = @file_get_contents($apiUrl);

        // Check if the HTTP request was successful
        if ($jsonData !== false) {
            // Decode the JSON data
            $data = json_decode($jsonData, true);

            // Check if the decoding was successful
            if ($data !== null) {
                // Store the data in the session
                $_SESSION['user_data'] = $data;

                // Redirect to another page
                header("Location: show.php");
                exit();
            } else {
                echo "Failed to decode JSON data.";
            }
        } else {
            echo "Failed to retrieve JSON data. HTTP request failed.";
        }
    } else {
        echo "Error: Variable \$result is not set.";
    }
}

// Redirect back to sum.php if there was an issue
header("Location: sum.php");
exit();
?>
