package text

import (
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func genRandomIntsUnique(n, textLen int, seed int64) []int {
	rng := rand.New(rand.NewSource(seed))
	indices := make(map[int]bool)
	for range n {
		var r int
		for {
			r = rng.Intn(textLen)
			if indices[r] {
				continue
			}
			break
		}

		indices[r] = true
	}

	var randomSample []int
	for i := range indices {
		randomSample = append(randomSample, i)
	}

	return randomSample
}

func sliceIsUnique[T comparable](slice []T) bool {
	uniqueEls := make(map[T]bool)
	for _, el := range slice {
		if uniqueEls[el] {
			return false
		}
		uniqueEls[el] = true
	}
	return true
}

func TestGenRandomIntsUnique(t *testing.T) {
	seed := int64(33)
	sampleSize := 10
	textLen := len(strings.TrimSpace(dummyText))
	got := genRandomIntsUnique(sampleSize, textLen, seed)

	t.Run("Testing number of slice elements", func(t *testing.T) {
		gotLen := len(got)
		if gotLen != sampleSize {
			t.Errorf("Got %d elements, want %d elements", gotLen, sampleSize)
		}
	})
	t.Run("Testing that elements in slice are unique", func(t *testing.T) {
		if !sliceIsUnique(got) {
			t.Errorf("slice contains duplicate elements")
		}
	})
}

func getRandomSample(n int, seed int64) []string {
	textSplit := strings.Fields(strings.TrimSpace(dummyText))
	textLen := len(textSplit)

	var textSample []string
	for i := range genRandomIntsUnique(n, textLen, seed) {
		textSample = append(textSample, strings.ToLower(textSplit[i]))
	}

	return textSample
}

func TestShuffleText(t *testing.T) {
	got := shuffleText(dummyText)

	t.Run("String length", func(t *testing.T) {
		gotLen := len(got)
		wantLen := len(strings.TrimSpace(dummyText))
		if gotLen != wantLen {
			t.Errorf("Got length %d, want length %d", gotLen, wantLen)
		}
	})

	t.Run("Parts of original string are in shuffled string", func(t *testing.T) {
		textSample := getRandomSample(10, 33)
		for _, word := range textSample {
			if !strings.Contains(strings.ToLower(strings.TrimSpace(dummyText)), word) {
				t.Error("Shuffled string does not contain all words in sample taken from original string")
			}
		}
	})
}

func TestGenText(t *testing.T) {
	type args struct {
		numParagraphs int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Paragraphs: 1", args{1}, 1},
		{"Paragraphs: 3", args{3}, 3},
		{"Paragraphs: 5", args{5}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := len(genText(tt.args.numParagraphs))
			if got != tt.want {
				t.Errorf("genText() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestTextGenDummyFiles(t *testing.T) {
	// Create a temp dir that will be used as the root for all tests
	dir, err := os.MkdirTemp("", "textgen-test")
	if err != nil {
		t.Fatalf("Could not create temp directory: %s", err)
	}
	defer os.RemoveAll(dir)

	// Run the tests
	t.Run("Testing files", testGenDummyFiles_Files(dir))
	t.Run("Testing directories", testGenDummyFiles_Directories(dir))
}

func testGenDummyFiles_Directories(tempDirPath string) func(t *testing.T) {
	return func(t *testing.T) {
		type args struct {
			outputPath string
		}
		tests := []struct {
			name string
			args args
		}{
			{"Num dirs: 0", args{""}},
			{"Num dirs: 1", args{"out"}},
			{"Num dirs: 3", args{"test/one/two"}},
		}

		for i, tt := range tests {
			testDirName := fmt.Sprintf("test_dirs%d", i+1)
			testDir := filepath.Join(tempDirPath, testDirName)
			err := os.Mkdir(testDir, 0755)
			if err != nil {
				t.Fatalf("Could not create temp directory: %s", err)
			}
			defer os.RemoveAll(testDir)

			t.Run(tt.name, func(t *testing.T) {
				outputDir := filepath.Join(testDir, tt.args.outputPath)
				err := GenDummyFiles(1, 1, outputDir)
				if err != nil {
					t.Fatalf("GenDummyFiles ran into an error: %v", err)
				}

				var directories []string
				err = filepath.WalkDir(testDir, func(path string, d fs.DirEntry, err error) error {
					if err != nil {
						return err
					}
					if d.IsDir() && d.Name() != testDirName {
						directories = append(directories, d.Name())
					}
					return nil
				})
				if err != nil {
					t.Fatalf("Problem walking directories: %v", err)
				}

				var got string
				if len(directories) > 0 {
					got = strings.Join(directories, "/")
				} else {
					got = ""
				}

				want := tt.args.outputPath
				if got != want {
					t.Fatalf("Got '%s', want '%s'", got, want)
				}
			})
		}
	}
}

func testGenDummyFiles_Files(tempDirPath string) func(t *testing.T) {
	return func(t *testing.T) {
		type args struct {
			numFiles int
		}
		tests := []struct {
			name string
			args args
		}{
			{"Num files: 1", args{1}},
			{"Num files: 3", args{3}},
			{"Num files: 5", args{5}},
		}

		for i, tt := range tests {
			testDirName := fmt.Sprintf("test_files%d", i+1)
			testDir := filepath.Join(tempDirPath, testDirName)
			err := os.Mkdir(testDir, 0755)
			if err != nil {
				t.Fatalf("Could not create temp directory: %s", err)
			}
			defer os.RemoveAll(testDir)

			t.Run(tt.name, func(t *testing.T) {
				err := GenDummyFiles(tt.args.numFiles, 1, testDir)
				if err != nil {
					t.Fatalf("GenDummyFiles ran into an error: %v", err)
				}

				files, err := os.ReadDir(testDir)
				if err != nil {
					t.Fatalf("Problem reading files in '%s': %v", testDir, err)
				}

				got := len(files)
				want := tt.args.numFiles
				if got != want {
					t.Fatalf("Got %d files, want %d files", got, want)
				}
			})
		}
	}
}
