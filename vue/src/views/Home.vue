<template>
  <div>
    <div class="logo-text">
      <p>Customer</p>
      <p>Management</p>
      <p>Test</p>
    </div>
  </div>
  <div>
    <CustomerList />
  </div>
</template>

<script>
import CustomerList from "../components/CustomerList.vue";

export default {
  components: {
    CustomerList,
  },
};
</script>

<style>
p {
  margin: 0;
}

.logo-text {
  border: 1px solid gray;
  border-radius: 4px;
  padding: 8px;
  text-align: center;
  width: fit-content;
}
</style>
