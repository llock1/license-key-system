
export function validatePassword(password) {
    const minLength = 8;
    const maxLength = 64;
    const hasSpecialChar = /[!@#$%^&*(),.?":{}|<>]/.test(password);

    if (password.length >= minLength && password.length <= maxLength) {
        return true;
    } else {
        return false;
    }
}
