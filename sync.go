// Copyright (C) 2018  denis4net
// Copyright (C) 2019 Evgeny Kuznetsov (evgeny@kuznetsov.md)
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

   photoslibrary "github.com/carstentrink/gphotoslibrary"
   "github.com/carstentrink/gphotosync/internal/oauth"
   "google.golang.org/api/googleapi"
)

type Library struct {
	Path         string
	Deduplicator func(*photoslibrary.MediaItem) string
}

func DownloadFile(url string, path string) error {

	// Create the file
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func getDownloadUrl(mediaItem *photoslibrary.MediaItem) string {
	if mediaItem.MediaMetadata.Photo == nil {
		return fmt.Sprintf("%s=dv", mediaItem.BaseUrl)
	}

	return fmt.Sprintf("%s=d", mediaItem.BaseUrl)

}

func (lib *Library) GetTokenPath() string {
	return path.Join(lib.Path, "USERNAME.token.json")
}


func (lib *Library) SyncMediaItem(mItem *photoslibrary.MediaItem) error {
	remoteCreationTime, err := time.Parse(time.RFC3339, mItem.MediaMetadata.CreationTime)
	if err != nil {
		return err
	}

	mediaPath, err := getMediaPath(lib, mItem)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(path.Dir(mediaPath), 0755); err != nil {
		return err
	}

	// multiple items with same filename can exist in remote
	// we try to mitigate it by adding timestamp to local filename
	if stat, err := os.Stat(mediaPath); !os.IsNotExist(err) && stat.ModTime().UnixNano() != remoteCreationTime.UnixNano() {
//		mediaPath, err = deduplicatePath(lib, mItem)
		if err != nil {
			return err
		}
	}

	if _, err := os.Stat(mediaPath); os.IsNotExist(err) {
		mediaURL := getDownloadUrl(mItem)


        searchText1 := "ABCDE"
        searchText2 := "ABCDE"
        searchText3 := "ABCDE"
        searchText4 := "ABCDE"
        searchText5 := "ABCDE"
        searchText6 := "ABCDE"
        searchText7 := "ABCDE"
        searchText8 := "ABCDE"
        searchText9 := "ABCDE"
//  current_time := time.Now().Local()
//  fmt.Println("The Current time is ", time.Now().Local().Unix().Format("2006-01-02"))
//  fmt.Println("File time is ", remoteCreationTime.Format("2006-01-02"))
//  timestamp_now := time.Now().Local().Unix()
//  filetimestamp := remoteCreationTime.Unix()
//  fmt.Printf("unix NOW \"%d\" but file is \"%d\" \n", timestamp_now,filetimestamp)
//  diff := timestamp_now - filetimestamp
//  if diff > 10 {
//  if diff > 99999999259200 {
//  fmt.Printf("checking if newer than 48 hours \"%d\" \n", diff)


//  fmt.Println("Unix %s---- ", timestamp)
//        t2 := t.Add(time.Minute * 3600)
//        fmt.Printf("CURRENT %s  \n", start)

        if strings.Contains(mediaPath, searchText1) {
            fmt.Printf("Downloading \"%s\" \n", mediaPath)
            err := DownloadFile(mediaURL, mediaPath)
            if err != nil {
                return err
            }

            err = os.Chtimes(mediaPath, remoteCreationTime, remoteCreationTime)
            if err != nil {
                return err
            }
        } else if strings.Contains(mediaPath, searchText2) {
            fmt.Printf("Downloading \"%s\" \n", mediaPath)
            err := DownloadFile(mediaURL, mediaPath)
            if err != nil {
                return err
            }

            err = os.Chtimes(mediaPath, remoteCreationTime, remoteCreationTime)
            if err != nil {
                return err
            }
        } else if strings.Contains(mediaPath, searchText3) {
            fmt.Printf("Downloading \"%s\" \n", mediaPath)
            err := DownloadFile(mediaURL, mediaPath)
            if err != nil {
                return err
            }

            err = os.Chtimes(mediaPath, remoteCreationTime, remoteCreationTime)
            if err != nil {
                return err
            }
        } else if strings.Contains(mediaPath, searchText4) {
            fmt.Printf("Downloading \"%s\" \n", mediaPath)
            err := DownloadFile(mediaURL, mediaPath)
            if err != nil {
                return err
            }

            err = os.Chtimes(mediaPath, remoteCreationTime, remoteCreationTime)
            if err != nil {
                return err
            }
        } else if strings.Contains(mediaPath, searchText5) {
            fmt.Printf("Downloading \"%s\" \n", mediaPath)
            err := DownloadFile(mediaURL, mediaPath)
            if err != nil {
                return err
            }
            err = os.Chtimes(mediaPath, remoteCreationTime, remoteCreationTime)
            if err != nil {
                return err
            }
        } else if strings.Contains(mediaPath, searchText6) {
            fmt.Printf("Downloading \"%s\" \n", mediaPath)
            err := DownloadFile(mediaURL, mediaPath)
            if err != nil {
                return err
            }

            err = os.Chtimes(mediaPath, remoteCreationTime, remoteCreationTime)
            if err != nil {
                return err
            }
        } else if strings.Contains(mediaPath, searchText7) {
            fmt.Printf("Downloading \"%s\" \n", mediaPath)
            err := DownloadFile(mediaURL, mediaPath)
            if err != nil {
                return err
            }

            err = os.Chtimes(mediaPath, remoteCreationTime, remoteCreationTime)
            if err != nil {
                return err
            }
        } else if strings.Contains(mediaPath, searchText8) {
            fmt.Printf("Downloading \"%s\" \n", mediaPath)
            err := DownloadFile(mediaURL, mediaPath)
            if err != nil {
                return err
            }

            err = os.Chtimes(mediaPath, remoteCreationTime, remoteCreationTime)
            if err != nil {
                return err
            }
        } else if strings.Contains(mediaPath, searchText9) {
            fmt.Printf("Downloading \"%s\" \n", mediaPath)
            err := DownloadFile(mediaURL, mediaPath)
            if err != nil {
                return err
            }

            err = os.Chtimes(mediaPath, remoteCreationTime, remoteCreationTime)
            if err != nil {
                return err
            }
        } else {
        //fmt.Printf("%s is ignored %s\n", mItem.Filename, mItem.MediaMetadata.CreationTime)
        //fmt.Printf("LIST %s\n", mItem.Filename)
        }
//	}
	} else {
		fmt.Printf("\"%s\"\n", mediaPath)
	}

	return nil
}

func getMediaPath(lib *Library, mItem *photoslibrary.MediaItem) (string, error) {
	//remoteCreationTime, err := time.Parse(time.RFC3339, mItem.MediaMetadata.CreationTime)
	//if err != nil {
	//	return "", err
	//}
	//return path.Join(lib.Path, strconv.Itoa(remoteCreationTime.Year()), fmt.Sprintf("%02d", remoteCreationTime.Month()), mItem.Filename), nil
	return (mItem.Filename), nil
}

//func deduplicatePath(lib *Library, item *photoslibrary.MediaItem) (string, error) {
//	p, err := getMediaPath(lib, item)
//	if err != nil {
//		return "", err
//	}
//	e := path.Ext(p)
//	return strings.TrimSuffix(p, e) + "-gphotosync-" + lib.Deduplicator(item) + e, nil
//}

func dedupUnixHex(item *photoslibrary.MediaItem) string {
	remoteCreationTime, err := time.Parse(time.RFC3339, item.MediaMetadata.CreationTime)
	if err != nil {
		return ""
	}
	return strconv.FormatInt(remoteCreationTime.UnixNano(), 16)
}

func dedupID(item *photoslibrary.MediaItem) string {
	return item.Id
}

func (lib *Library) Sync(cred credentials) error {
	ctx := context.Background()

	oauthClient, err := oauth.NewClient(ctx, cred.ID, cred.Secret, cred.RedirectURL, lib.GetTokenPath())
	if err != nil {
		return err
	}
	photoslibraryService, err := photoslibrary.New(oauthClient)
	if err != nil {
		return err
	}

	mediaItemsService := photoslibrary.NewMediaItemsService(photoslibraryService)
	pageToken := ""
	waitTime := 1 * time.Minute

	for {
		res, err := mediaItemsService.Search(&photoslibrary.SearchMediaItemsRequest{PageToken: pageToken, PageSize: 100}).Do()
		if err != nil {
			if apiError, ok := err.(*googleapi.Error); ok {
				/// If the quota of requests to the Library API is exceeded, the API returns an error code 429 and a message that the project has exceeded the quota.
				switch apiError.Code {
				case 429:
					fmt.Printf("failed to get media items: %s\nwill wait and try again\n", apiError.Message)
					time.Sleep(waitTime)
					waitTime = waitTime * 2
					continue
				case 500:
					fmt.Println("got error 500 from server, let's wait and see")
					time.Sleep(time.Minute)
					continue
				case 502:
					fmt.Println("got error 502 from Google API, will retry in 30 seconds")
					time.Sleep(30 * time.Second)
					continue
				case 503:
					fmt.Println("got error 503 from server, will wait for a minute and try again")
					time.Sleep(time.Minute)
					continue
				default:
					return err
				}
			}
		}

		// sometimes the library returns both err and res empty (yep, a bug in there)
		if res == nil {
			fmt.Println("looks like Google returned empty page...")
			return nil
		} else {
			fmt.Printf("processing %d items\n", len(res.MediaItems))
			for _, mItem := range res.MediaItems {
				err = lib.SyncMediaItem(mItem)
				if err != nil {
					log.Printf("failed to sync item \"%s\": %s", mItem.Id, err)
				}
			}
		}

		// if NextPageToken is empty, we reached the last page of items list
		if res.NextPageToken == "" {
			//fmt.Println("syncing is done")
			return nil
		} else {
			fmt.Println("requesting next page")
		}

		pageToken = res.NextPageToken
		waitTime = 1 * time.Minute
	}
}
