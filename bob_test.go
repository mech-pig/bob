package bob

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type testTargetStruct struct {
	a string
	b string
}

func TestBuildDefault(t *testing.T) {
	defaultInstance := testTargetStruct{a: "a", b: "b"}
	factory := New(func() testTargetStruct { return defaultInstance })
	require.Equal(t, defaultInstance, factory.Build())
}

func TestBuildWithOverrides(t *testing.T) {
	defaultInstance := testTargetStruct{a: "a", b: "b"}
	factory := New(func() testTargetStruct { return defaultInstance })

	testCases := []struct {
		name      string
		overrides []func(testTargetStruct) testTargetStruct
		expected  testTargetStruct
	}{
		{
			name:      "nil",
			overrides: nil,
			expected:  defaultInstance,
		},
		{
			name:      "empty",
			overrides: []func(testTargetStruct) testTargetStruct{},
			expected:  defaultInstance,
		},
		{
			name: "single override",
			overrides: []func(testTargetStruct) testTargetStruct{
				func(i testTargetStruct) testTargetStruct {
					i.a = "A"
					return i
				},
			},
			expected: testTargetStruct{
				a: "A",
				b: defaultInstance.b,
			},
		},
		{
			name: "multiple overrides",
			overrides: []func(testTargetStruct) testTargetStruct{
				func(i testTargetStruct) testTargetStruct {
					i.a = "A"
					return i
				},
				func(i testTargetStruct) testTargetStruct {
					i.b = "B"
					return i
				},
			},
			expected: testTargetStruct{
				a: "A",
				b: "B",
			},
		},
		{
			name: "overrides are applied from left to right",
			overrides: []func(testTargetStruct) testTargetStruct{
				func(i testTargetStruct) testTargetStruct {
					i.a = "first"
					return i
				},
				func(i testTargetStruct) testTargetStruct {
					i.a = "second"
					return i
				},
			},
			expected: testTargetStruct{
				a: "second",
				b: defaultInstance.b,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, factory.Build(tc.overrides...))
		})
	}
}

func TestBuildManyDefault(t *testing.T) {
	defaultInstance := testTargetStruct{a: "a", b: "b"}
	factory := New(func() testTargetStruct { return defaultInstance })
	expected := []testTargetStruct{defaultInstance, defaultInstance, defaultInstance}
	require.Equal(t, expected, factory.BuildMany(len(expected)))
}

func TestBuildManyWithOverrides(t *testing.T) {
	defaultInstance := testTargetStruct{a: "a", b: "b"}
	factory := New(func() testTargetStruct { return defaultInstance })

	testCases := []struct {
		name      string
		overrides []func(int, testTargetStruct) testTargetStruct
		expected  []testTargetStruct
	}{
		{
			name:      "nil",
			overrides: nil,
			expected:  []testTargetStruct{defaultInstance, defaultInstance},
		},
		{
			name:      "empty",
			overrides: []func(int, testTargetStruct) testTargetStruct{},
			expected:  []testTargetStruct{defaultInstance, defaultInstance},
		},
		{
			name: "single override",
			overrides: []func(int, testTargetStruct) testTargetStruct{
				func(i int, instance testTargetStruct) testTargetStruct {
					instance.a = fmt.Sprint(i)
					return instance
				},
			},
			expected: []testTargetStruct{
				{
					a: "0",
					b: defaultInstance.b,
				},
				{
					a: "1",
					b: defaultInstance.b,
				},
			},
		},
		{
			name: "multiple overrides",
			overrides: []func(int, testTargetStruct) testTargetStruct{
				func(i int, instance testTargetStruct) testTargetStruct {
					instance.a = fmt.Sprint(i)
					return instance
				},
				func(i int, instance testTargetStruct) testTargetStruct {
					instance.b = fmt.Sprint(i)
					return instance
				},
			},
			expected: []testTargetStruct{
				{
					a: "0",
					b: "0",
				},
				{
					a: "1",
					b: "1",
				},
			},
		},

		{
			name: "overrides are applied from left to right",
			overrides: []func(int, testTargetStruct) testTargetStruct{
				func(_ int, instance testTargetStruct) testTargetStruct {
					instance.a = "first"
					return instance
				},
				func(_ int, instance testTargetStruct) testTargetStruct {
					instance.a = "second"
					return instance
				},
			},
			expected: []testTargetStruct{
				{
					a: "second",
					b: defaultInstance.b,
				},
				{
					a: "second",
					b: defaultInstance.b,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, factory.BuildMany(len(tc.expected), tc.overrides...))
		})
	}
}

func TestOverride(t *testing.T) {
	defaultInstance := testTargetStruct{a: "a", b: "b"}
	factory := New(func() testTargetStruct { return defaultInstance })
	derived := factory.Override(func(instance testTargetStruct) testTargetStruct {
		instance.a = "derived"
		return instance
	})

	testCases := []struct {
		name     string
		actual   any
		expected any
	}{
		{
			name:   "build",
			actual: derived.Build(),
			expected: testTargetStruct{
				a: "derived",
				b: defaultInstance.b,
			},
		},
		{
			name: "build with overrides",
			actual: derived.Build(func(instance testTargetStruct) testTargetStruct {
				instance.b = "overridden"
				return instance
			}),
			expected: testTargetStruct{
				a: "derived",
				b: "overridden",
			},
		},
		{
			name:   "build many",
			actual: derived.BuildMany(2),
			expected: []testTargetStruct{
				{
					a: "derived",
					b: defaultInstance.b,
				},
				{
					a: "derived",
					b: defaultInstance.b,
				},
			},
		},
		{
			name: "build many with overrides",
			actual: derived.BuildMany(2, func(i int, instance testTargetStruct) testTargetStruct {
				instance.b = fmt.Sprint("overridden-", i)
				return instance
			}),
			expected: []testTargetStruct{
				{
					a: "derived",
					b: "overridden-0",
				},
				{
					a: "derived",
					b: "overridden-1",
				},
			},
		},
		{
			name: "override",
			actual: derived.
				Override(func(instance testTargetStruct) testTargetStruct {
					instance.b = "derived"
					return instance
				}).
				Build(),
			expected: testTargetStruct{
				a: "derived",
				b: "derived",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, tc.actual)
		})
	}
}
