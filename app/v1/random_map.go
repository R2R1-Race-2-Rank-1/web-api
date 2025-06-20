package v1

import (
    "database/sql"

    "zxq.co/ripple/rippleapi/common"
)

// randomBeatmapSelect is a SQL fragment that selects the typical beatmap fields.
const randomBeatmapSelect = `
SELECT beatmap_id, beatmapset_id, beatmap_md5,
       song_name, ar, od, difficulty_std, difficulty_taiko,
       difficulty_ctb, difficulty_mania, max_combo, hit_length,
       ranked, ranked_status_freezed, latest_update
FROM beatmaps `

// RandomMapGET retrieves a random beatmap from the database.
func RandomMapGET(md common.MethodData) common.CodeMessager {
    var b beatmap
    // Append ORDER BY RAND() to randomly select one row.
    query := randomBeatmapSelect + " ORDER BY RAND() LIMIT 1"

    err := md.DB.QueryRow(query).Scan(
        &b.BeatmapID,
        &b.BeatmapsetID,
        &b.BeatmapMD5,
        &b.SongName,
        &b.AR,
        &b.OD,
        &b.Diff2.STD,
        &b.Diff2.Taiko,
        &b.Diff2.CTB,
        &b.Diff2.Mania,
        &b.MaxCombo,
        &b.HitLength,
        &b.Ranked,
        &b.RankedStatusFrozen,
        &b.LatestUpdate,
    )
    if err != nil {
        if err == sql.ErrNoRows {
            return common.SimpleResponse(404, "No beatmap found")
        }
        md.Err(err)
        return common.SimpleResponse(500, "Database error")
    }

    return struct {
        common.ResponseBase
        Map beatmap `json:"map"`
    }{
        ResponseBase: common.ResponseBase{Code: 200},
        Map:          b,
    }
}