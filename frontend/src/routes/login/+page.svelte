<script>
  function redirectToGoogleLogin() {
    window.location.href = 'http://localhost:8080/auth/google/login';
  }
</script>

<div class="container">
  <h1>Login</h1>
  <p>Please log in using your Google account.</p>
  <hr>

  <button on:click={redirectToGoogleLogin} class="google-login-btn">
    Login with Google
  </button>

  <div class="container login">
    <p>Don't have an account? Register by logging in with Google.</p>
  </div>
</div>

<style>
  .container {
    max-width: 400px;
    margin: 50px auto;
    padding: 20px;
    border: 1px solid #ddd;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    text-align: center;
  }

  h1 {
    color: #333;
    margin-bottom: 20px;
  }

  p {
    color: #666;
    margin-bottom: 20px;
  }

  hr {
    border: none;
    border-top: 1px solid #eee;
    margin: 20px 0;
  }

  .google-login-btn {
    background-color: #4285F4;
    color: white;
    padding: 12px 20px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    font-size: 16px;
    transition: background-color 0.3s ease;
  }

  .google-login-btn:hover {
    background-color: #357ae8;
  }

  .login p {
    margin-top: 20px;
    font-size: 14px;
  }

  .login a {
    color: #007bff;
    text-decoration: none;
  }

  .login a:hover {
    text-decoration: underline;
  }
</style>