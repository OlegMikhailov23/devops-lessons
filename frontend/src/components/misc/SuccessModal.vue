<template>
  <div v-if="show" class="modal fade show" tabindex="-1" style="display: block; background-color: rgba(0,0,0,0.5);">
    <div class="modal-dialog modal-dialog-centered">
      <div class="modal-content">
        <div class="modal-body text-center py-4">
          <div class="mb-3">
            <i class="bi bi-check-circle-fill text-success" style="font-size: 3rem;"></i>
          </div>
          <h5 class="modal-title">{{ title }}</h5>
          <p class="mb-0">{{ message }}</p>
        </div>
        <div class="modal-footer justify-content-center">
          <button type="button" class="btn btn-success" @click="close">
            Отлично!
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'SuccessModal',
  props: {
    show: {
      type: Boolean,
      default: false
    },
    title: {
      type: String,
      default: 'Успешно!'
    },
    message: {
      type: String,
      default: ''
    }
  },
  emits: ['close'],
  methods: {
    close() {
      this.$emit('close')
    }
  }
})
</script>

<style scoped>
.modal.show {
  background-color: rgba(0,0,0,0.5);
}

.modal-dialog {
  max-width: 400px;
}

.bi-check-circle-fill {
  color: #28a745 !important;
}
</style>
