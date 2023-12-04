<script lang="ts">
  import ThemeChanger from './ThemeChanger.svelte';
</script>

<header class="navbar bg-base-100">
  <div class="navbar-start">
    <div class="dropdown">
      <div tabindex="0" role="button" class="btn btn-ghost btn-circle">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-5 w-5"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
          ><path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M4 6h16M4 12h16M4 18h7"
          /></svg
        >
      </div>
      <ul
        class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
      >
        <li>Settings</li>
      </ul>
    </div>
  </div>
  <div class="navbar-center">
    <p class="btn btn-ghost text-xl">Mayoi 🐌</p>
  </div>
  <div class="navbar-end">
    <ThemeChanger />
  </div>
</header>
