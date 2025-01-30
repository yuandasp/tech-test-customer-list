<template>
  <div class="form">
    <h2 style="text-align: center; color: rgb(84, 84, 238)">
      {{ editData ? "Edit Customer" : "Add New Customer" }}
    </h2>
    <form @submit.prevent="submitForm">
      <div class="container-inside-form">
        <div class="box-input">
          <div style="width: 25%">
            <label>Name</label>
          </div>
          <div style="width: 70%">
            <input class="input-style" v-model="form.name" required />
          </div>
        </div>

        <div class="box-input">
          <div style="width: 25%">
            <label>Birth Date</label>
          </div>
          <div style="width: 70%">
            <input
              class="input-style"
              type="date"
              v-model="form.birth_date"
              required
            />
          </div>
        </div>

        <div class="box-input">
          <div style="width: 25%">
            <label>Address</label>
          </div>

          <div style="width: 70%">
            <input class="input-style" v-model="form.address" required />
          </div>
        </div>

        <div class="box-input">
          <div style="width: 25%">
            <label>Phone Number</label>
          </div>

          <div style="width: 70%">
            <input class="input-style" v-model="form.phone_number" required />
          </div>
        </div>
      </div>
      <button class="btn" type="submit">
        {{ editData ? "Update" : "Create" }}
      </button>
    </form>
  </div>
</template>

<script>
import axios from "axios";

export default {
  props: ["editData"],
  data() {
    return {
      form: {
        name: "",
        birth_date: "",
        address: "",
        phone_number: "",
      },
    };
  },
  watch: {
    editData: {
      immediate: true,
      handler(newVal) {
        if (newVal) {
          this.form = { ...newVal };
        } else {
          this.resetForm();
        }
      },
    },
  },
  methods: {
    async submitForm() {
      try {
        if (this.editData) {
          // Update customer
          await axios.put(
            `${import.meta.env.VITE_API_URL}/customers/${this.editData.id}`,
            this.form
          );
        } else {
          // Create customer
          await axios.post(
            `${import.meta.env.VITE_API_URL}/customers/`,
            this.form
          );
        }
        this.$emit("customer-updated");
        this.resetForm();
      } catch (error) {
        console.error("Error submitting form:", error);
      }
    },
    resetForm() {
      this.form = {
        name: "",
        birth_date: "",
        address: "",
        phone_number: "",
      };
    },
  },
};
</script>

<style>
p {
  margin: 0;
}

.form {
  margin-top: 24px;
  padding: 16px;
  border-radius: 8px;
  width: 40%;
}

.container-inside-form {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.box-input {
  display: flex;
  gap: 24px;
  justify-content: space-between;
  width: 100%;
}

.flex-input {
  display: flex;
  gap: 4px;
}

.input-style {
  border: 1px solid rgb(214, 214, 214);
  border-radius: 4px;
  padding: 8px 16px;
  width: 90%;
}

.btn {
  width: 100%;
  background-color: rgb(131, 131, 244);
  border: 1px solid rgb(131, 131, 244);
  margin-top: 16px;
  border-radius: 8px;
  padding: 8px;
  color: #fff;
  font-weight: 600;
  font-size: 16px;
}

.btn:hover {
  background-color: rgb(153, 153, 244);
  border: 1px solid rgb(153, 153, 244);
  cursor: pointer;
}
</style>
