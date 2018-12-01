// Write your event here

package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/ying32/govcl/vcl/types"

	"github.com/ying32/govcl/vcl"
)

//::private::
type TMainFormFields struct {
}

const jsonData = `{
  "id": 104690469,
  "node_id": "MDEwOlJlcG9zaXRvcnkxMDQ2OTA0Njk=",
  "name": "govcl",
  "full_name": "ying32/govcl",
  "private": false,
  "owner": {
    "login": "ying32",
    "id": 7037165,
    "node_id": "MDQ6VXNlcjcwMzcxNjU=",
    "avatar_url": "https://avatars1.githubusercontent.com/u/7037165?v=4",
    "gravatar_id": "",
    "url": "https://api.github.com/users/ying32",
    "html_url": "https://github.com/ying32",
    "followers_url": "https://api.github.com/users/ying32/followers",
    "following_url": "https://api.github.com/users/ying32/following{/other_user}",
    "gists_url": "https://api.github.com/users/ying32/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/ying32/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/ying32/subscriptions",
    "organizations_url": "https://api.github.com/users/ying32/orgs",
    "repos_url": "https://api.github.com/users/ying32/repos",
    "events_url": "https://api.github.com/users/ying32/events{/privacy}",
    "received_events_url": "https://api.github.com/users/ying32/received_events",
    "type": "User",
    "site_admin": false
  },
  "html_url": "https://github.com/ying32/govcl",
  "description": " A cross-platform Golang GUI library. Use Delphi VCL and Lazarus LCL for binding. ",
  "fork": false,
  "url": "https://api.github.com/repos/ying32/govcl",
  "forks_url": "https://api.github.com/repos/ying32/govcl/forks",
  "keys_url": "https://api.github.com/repos/ying32/govcl/keys{/key_id}",
  "collaborators_url": "https://api.github.com/repos/ying32/govcl/collaborators{/collaborator}",
  "teams_url": "https://api.github.com/repos/ying32/govcl/teams",
  "hooks_url": "https://api.github.com/repos/ying32/govcl/hooks",
  "issue_events_url": "https://api.github.com/repos/ying32/govcl/issues/events{/number}",
  "events_url": "https://api.github.com/repos/ying32/govcl/events",
  "assignees_url": "https://api.github.com/repos/ying32/govcl/assignees{/user}",
  "branches_url": "https://api.github.com/repos/ying32/govcl/branches{/branch}",
  "tags_url": "https://api.github.com/repos/ying32/govcl/tags",
  "blobs_url": "https://api.github.com/repos/ying32/govcl/git/blobs{/sha}",
  "git_tags_url": "https://api.github.com/repos/ying32/govcl/git/tags{/sha}",
  "git_refs_url": "https://api.github.com/repos/ying32/govcl/git/refs{/sha}",
  "trees_url": "https://api.github.com/repos/ying32/govcl/git/trees{/sha}",
  "statuses_url": "https://api.github.com/repos/ying32/govcl/statuses/{sha}",
  "languages_url": "https://api.github.com/repos/ying32/govcl/languages",
  "stargazers_url": "https://api.github.com/repos/ying32/govcl/stargazers",
  "contributors_url": "https://api.github.com/repos/ying32/govcl/contributors",
  "subscribers_url": "https://api.github.com/repos/ying32/govcl/subscribers",
  "subscription_url": "https://api.github.com/repos/ying32/govcl/subscription",
  "commits_url": "https://api.github.com/repos/ying32/govcl/commits{/sha}",
  "git_commits_url": "https://api.github.com/repos/ying32/govcl/git/commits{/sha}",
  "comments_url": "https://api.github.com/repos/ying32/govcl/comments{/number}",
  "issue_comment_url": "https://api.github.com/repos/ying32/govcl/issues/comments{/number}",
  "contents_url": "https://api.github.com/repos/ying32/govcl/contents/{+path}",
  "compare_url": "https://api.github.com/repos/ying32/govcl/compare/{base}...{head}",
  "merges_url": "https://api.github.com/repos/ying32/govcl/merges",
  "archive_url": "https://api.github.com/repos/ying32/govcl/{archive_format}{/ref}",
  "downloads_url": "https://api.github.com/repos/ying32/govcl/downloads",
  "issues_url": "https://api.github.com/repos/ying32/govcl/issues{/number}",
  "pulls_url": "https://api.github.com/repos/ying32/govcl/pulls{/number}",
  "milestones_url": "https://api.github.com/repos/ying32/govcl/milestones{/number}",
  "notifications_url": "https://api.github.com/repos/ying32/govcl/notifications{?since,all,participating}",
  "labels_url": "https://api.github.com/repos/ying32/govcl/labels{/name}",
  "releases_url": "https://api.github.com/repos/ying32/govcl/releases{/id}",
  "deployments_url": "https://api.github.com/repos/ying32/govcl/deployments",
  "created_at": "2017-09-25T01:38:08Z",
  "updated_at": "2018-11-26T01:29:27Z",
  "pushed_at": "2018-11-30T06:15:04Z",
  "git_url": "git://github.com/ying32/govcl.git",
  "ssh_url": "git@github.com:ying32/govcl.git",
  "clone_url": "https://github.com/ying32/govcl.git",
  "svn_url": "https://github.com/ying32/govcl",
  "homepage": "",
  "size": 9103,
  "stargazers_count": 132,
  "watchers_count": 132,
  "language": "Go",
  "has_issues": true,
  "has_projects": true,
  "has_downloads": true,
  "has_wiki": true,
  "has_pages": false,
  "forks_count": 15,
  "mirror_url": null,
  "archived": false,
  "open_issues_count": 0,
  "license": {
    "key": "apache-2.0",
    "name": "Apache License 2.0",
    "spdx_id": "Apache-2.0",
    "url": "https://api.github.com/licenses/apache-2.0",
    "node_id": "MDc6TGljZW5zZTI="
  },
  "forks": 15,
  "open_issues": 0,
  "watchers": 132,
  "default_branch": "master",
  "network_count": 15,
  "subscribers_count": 20
}`

const jsonData2 = `{
   "key":1,
   "key2":[
   {"aaa":"aaaa", "a122":"a1"}, 
   {"bbb": "bbb"}
  ],
  "key3": [
     1, 2,3
   ],
     "key4": [
     "str1", "str2"
   ]
}
`

// https://api.github.com/repos/ying32/govcl

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {

}

func (f *TMainForm) jsonTree(str string) {
	var data interface{}
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		vcl.ShowMessage(err.Error())
		return
	}
	f.TreeView1.Items().BeginUpdate()
	defer f.TreeView1.Items().EndUpdate()
	root := f.TreeView1.Items().AddChild(nil, "JSON")
	f.buildTree(root, data, "")
	f.TreeView1.FullExpand()
}

func (f *TMainForm) buildTree(node *vcl.TTreeNode, data interface{}, keyName string) {

	switch data.(type) {
	case map[string]interface{}:

		var child *vcl.TTreeNode
		if keyName != "" {
			child = f.TreeView1.Items().AddChild(node, keyName)
			child = f.TreeView1.Items().AddChild(child, "Object")
		} else {
			child = f.TreeView1.Items().AddChild(node, "Object")
		}

		// 因为随机问题，考虑解决办法，有序总比乱序好。。。。
		keys := make([]string, 0)
		for key := range data.(map[string]interface{}) {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			if val, ok := data.(map[string]interface{})[key]; ok {
				f.buildTree(child, val, key)
			}
		}

	case []interface{}:
		var child *vcl.TTreeNode
		if keyName != "" {
			child = f.TreeView1.Items().AddChild(node, keyName)
			child = f.TreeView1.Items().AddChild(child, "Array")
		} else {
			child = f.TreeView1.Items().AddChild(node, "Array")
		}
		for _, val := range data.([]interface{}) {
			f.buildTree(child, val, "")
		}

	default:

		if node != nil && node.IsValid() {
			if keyName == "" {
				f.TreeView1.Items().AddChild(node, fmt.Sprintf("%v", data))
			} else {
				f.TreeView1.Items().AddChild(node, fmt.Sprintf("%s:%v", keyName, data))
			}
		}
	}
}

func (f *TMainForm) OnBtnPasteClick(sender vcl.IObject) {
	f.TreeView1.Items().Clear()
	if vcl.Clipboard.HasFormat(types.CF_TEXT) {
		f.jsonTree(vcl.Clipboard.AsText())
	}
}
