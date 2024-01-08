<?php
session_start();

// Retrieve the data from the session
$data = isset($_SESSION['user_data']) ? $_SESSION['user_data'] : [];

// Display the data
foreach ($data as $user) {
    echo "ID: " . $user['id'] . "<br>";
    echo "Username: " . $user['username'] . "<br>";
    echo "Email: " . $user['email'] . "<br>";
    echo "<hr>";
}

// Clear the session variable
unset($_SESSION['user_data']);
// header("Location: sum.php");
// exit();
?>
