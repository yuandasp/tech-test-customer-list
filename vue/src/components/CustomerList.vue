<template>
  <div class="layout">
    <div class="form-container">
      <FormCustomer @customer-updated="fetchCustomers" :edit-data="editData" />
    </div>

    <hr />

    <div class="table-container">
      <div class="box-table">
        <h2 style="text-align: center; color: rgb(84, 84, 238)">
          Customer List
        </h2>
        <table class="customer-table" border="1" cellpadding="8">
          <thead>
            <tr>
              <th>ID</th>
              <th>Name</th>
              <th>Birth Date</th>
              <th>Address</th>
              <th>Phone Number</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="customer in customers" :key="customer.id">
              <td>{{ customer.id }}</td>
              <td>{{ customer.name }}</td>
              <td>{{ customer.birth_date }}</td>
              <td>{{ customer.address }}</td>
              <td>{{ customer.phone_number }}</td>
              <td>
                <button class="btn-edit" @click="editCustomer(customer)">
                  Edit
                </button>
                <button class="btn-delete" @click="deleteCustomer(customer.id)">
                  Delete
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import FormCustomer from "./FormCustomer.vue";

export default {
  components: { FormCustomer },
  data() {
    return {
      customers: [],
      editData: null,
    };
  },
  methods: {
    async fetchCustomers() {
      try {
        const response = await axios.get(
          `${import.meta.env.VITE_API_URL}/customers/`
        );
        this.customers = response.data;
      } catch (error) {
        console.error("Error fetching customers:", error);
      }
    },
    editCustomer(customer) {
      this.editData = { ...customer };
    },
    async deleteCustomer(id) {
      try {
        await axios.delete(`${import.meta.env.VITE_API_URL}/customers/${id}`);
        this.fetchCustomers();
      } catch (error) {
        console.error("Error deleting customer:", error);
      }
    },
  },
  mounted() {
    this.fetchCustomers();
  },
};
</script>

<style>
.form-container {
  display: flex;
  justify-content: center;
}

.table-container {
  display: flex;
  justify-content: center;
}

.customer-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
  background: #fff;
  border-radius: 10px;
  overflow: hidden;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.customer-table thead {
  background-color: rgb(131, 131, 244);
  color: white;
  text-align: left;
}

.customer-table th,
.customer-table td {
  padding: 12px 15px;
  border: 1px solid #ddd;
}

.customer-table tbody tr {
  transition: background 0.3s;
}

.customer-table tbody tr:nth-child(even) {
  background-color: #f8f9fa;
}

.btn-delete {
  border-radius: 8px;
  background-color: rgb(218, 65, 65);
  border: 1px solid rgb(218, 65, 65);
  color: #fff;
  padding: 8px;
  font-weight: 600;
  margin-left: 4px;
}
.btn-edit {
  border-radius: 8px;
  background-color: rgb(131, 131, 244);
  border: 1px solid rgb(131, 131, 244);
  color: #fff;
  padding: 8px;
  font-weight: 600;
}

.btn-edit:hover {
  background-color: rgb(153, 153, 244);
  border: 1px solid rgb(153, 153, 244);
  cursor: pointer;
}

.btn-delete:hover {
  background-color: rgb(218, 86, 86);
  border: 1px solid rgb(218, 86, 86);
  cursor: pointer;
}
</style>
