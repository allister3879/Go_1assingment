document.addEventListener('DOMContentLoaded', function() {
    const emailInput = document.getElementById('email');
    const titleSelect = document.getElementById('title');
    const firstNameInput = document.getElementById('first-name');
    const lastNameInput = document.getElementById('last-name');
    // Add more element references as needed

    // Event Handling
    const submitButton = document.getElementById('submit-button');
    submitButton.addEventListener('click', function() {
        alert('Form submitted successfully!');
    });
});