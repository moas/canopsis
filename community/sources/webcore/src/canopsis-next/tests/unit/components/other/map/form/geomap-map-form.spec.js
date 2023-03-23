import Faker from 'faker';

import { mount, shallowMount, createVueInstance } from '@unit/utils/vue';

import GeomapMapForm from '@/components/other/map/form/geomap-map-form.vue';

const localVue = createVueInstance();

const stubs = {
  'c-name-field': true,
  'geomap-editor': true,
};

const factory = (options = {}) => shallowMount(GeomapMapForm, {
  localVue,
  stubs,

  ...options,
});

const snapshotFactory = (options = {}) => mount(GeomapMapForm, {
  localVue,
  stubs,

  ...options,
});

const selectNameField = wrapper => wrapper.find('c-name-field-stub');
const selectGeomapEditor = wrapper => wrapper.find('geomap-editor-stub');

describe('geomap-map-form', () => {
  test('Name changed after trigger name field', () => {
    const form = {
      name: '',
      parameters: {},
    };
    const wrapper = factory({
      propsData: {
        form,
      },
    });

    const newName = Faker.datatype.string();

    const nameField = selectNameField(wrapper);

    nameField.vm.$emit('input', newName);

    expect(wrapper).toEmit('input', {
      ...form,
      name: newName,
    });
  });

  test('Parameters changed after trigger mermaid editor field', () => {
    const form = {
      name: '',
      parameters: {},
    };
    const wrapper = factory({
      propsData: {
        form,
      },
    });

    const newParameters = {
      param: Faker.datatype.string(),
    };

    const geomapEditor = selectGeomapEditor(wrapper);

    geomapEditor.vm.$emit('input', newParameters);

    expect(wrapper).toEmit('input', {
      ...form,
      parameters: newParameters,
    });
  });

  test('Renders `geomap-map-form` with form', () => {
    const wrapper = snapshotFactory({
      propsData: {
        form: {
          name: 'Geomap',
          parameters: {},
        },
      },
    });

    expect(wrapper.element).toMatchSnapshot();
  });
});