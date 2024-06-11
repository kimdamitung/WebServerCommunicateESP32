function onlyNumbers(input) {
    var value = input.value;
    var sanitizedValue = value.replace(/[^0-9]/g, '');
    if (sanitizedValue.length > 10) {
        sanitizedValue = sanitizedValue.slice(0, 10);
    }
    input.value = sanitizedValue;
}

document.addEventListener('DOMContentLoaded', function () {
    document.getElementById('signup-sdt').addEventListener('input', function (e) {
        onlyNumbers(e.target);
    });
    document.getElementById('login-sdt').addEventListener('input', function (e) {
        onlyNumbers(e.target);
    });
});

const container = document.getElementById('container');
const registerBtn = document.getElementById('register');
const loginBtn = document.getElementById('login');

registerBtn.addEventListener('click', () => {
    container.classList.add("active");
});

loginBtn.addEventListener('click', () => {
    container.classList.remove("active");
});