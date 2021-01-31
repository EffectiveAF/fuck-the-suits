<script>
</script>

<div class="divider"></div>

<footer>
  <div class="logo-text">
    <p>Designed and Engineered in California</p>
  </div>

  <ul>
    <li>About us</li>
    <li>Vision</li>
  </ul>
</footer>

<style>
  .divider {
    width: calc(100vw - 16px);
    height: 1px;
    margin-top: 180px;
    background-color: var(--neutral200);
  }

  .logo-text {
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: center;
    width: 80%;
  }

  footer {
    display: flex;
    flex-direction: row;
    width: 100%;
    height: 64px;
  }

  footer > ul {
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
    align-items: center;
    width: 100%;
    height: 64px;
    list-style-type: none;
  }

  footer ul li {
    margin-left: 24px;
    text-decoration: none;
  }

  footer > * {
    padding: 0 24px;
  }

  @media (max-width: 499px) {
    .logo-text {
      width: 20%;
    }

    .logo-text > p {
      display: none;
    }

    footer > * {
      padding: 0 24px;
    }
  }
</style>
