<template>
  <div class="box">
    <div class="columns">
      <div class="column is-1">
        <span class="tag is-large is-rounded is-info">{{ index + 1 }}</span>
      </div>
      <div class="column">
        <div class="field is-horizontal">
          <div class="field-label is-normal">Segment:</div>
          <div class="field-body">
            <div class="field">
              <span class="tag is-medium"> {{ segmentName }} </span>
            </div>
          </div>
        </div>
        <div class="field is-horizontal">
          <div class="field-label is-normal">Then serve:</div>
          <div class="field-body">
            <div class="field">
              <div class="select">
                <select disabled>
                  <option v-if="distributions.length === 1">
                    {{ distributions[0].variantKey }}
                  </option>
                  <option v-else> A Percentage Rollout </option>
                </select>
              </div>
            </div>
          </div>
        </div>
        <template v-if="distributions.length > 1">
          <hr />
          <div
            v-for="(distribution, i) in distributions"
            :key="distribution.id"
            class="field is-horizontal"
          >
            <div class="field-label">
              <span class="tag is-small">{{ distribution.variantKey }}</span>
            </div>
            <div class="field-body">
              <div class="field">
                <BInput
                  v-model="distributions[i].rollout"
                  placeholder="Percentage"
                  disabled
                  type="number"
                  icon-pack="fas"
                  icon="percent"
                  size="is-small"
                  min="0"
                  max="100"
                />
              </div>
            </div>
          </div>
        </template>
      </div>
    </div>
    <div class="buttons is-right">
      <a
        v-if="distributions.length > 1"
        class="button is-text is-small"
        @click="editRule"
        >Edit</a
      >
      <a class="button is-danger is-outlined is-small" @click="deleteRule"
        >Delete</a
      >
    </div>
  </div>
</template>

<script>
export default {
  props: {
    id: String,
    segmentName: String,
    index: Number,
    distributions: Array
  },
  methods: {
    editRule() {
      this.$emit("editRule", this.index);
    },
    deleteRule() {
      this.$emit("deleteRule", this.index);
    }
  }
};
</script>
